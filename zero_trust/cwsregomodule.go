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

// CWSRegoModuleService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCWSRegoModuleService] method instead.
type CWSRegoModuleService struct {
	Options []option.RequestOption
}

// NewCWSRegoModuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCWSRegoModuleService(opts ...option.RequestOption) (r *CWSRegoModuleService) {
	r = &CWSRegoModuleService{}
	r.Options = opts
	return
}

// Creates a library Rego module and derives its package path from the source.
func (r *CWSRegoModuleService) New(ctx context.Context, params CWSRegoModuleNewParams, opts ...option.RequestOption) (res *RegoModule, err error) {
	var env CWSRegoModuleNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/rego-modules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the module and validates policies and libraries that depend on it.
func (r *CWSRegoModuleService) Update(ctx context.Context, regoModuleID string, params CWSRegoModuleUpdateParams, opts ...option.RequestOption) (res *RegoModule, err error) {
	var env CWSRegoModuleUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if regoModuleID == "" {
		err = errors.New("missing required rego_module_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/rego-modules/%s", params.AccountID, regoModuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of library Rego modules in the account.
func (r *CWSRegoModuleService) List(ctx context.Context, params CWSRegoModuleListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[RegoModule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/rego-modules", params.AccountID)
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

// Returns a paginated list of library Rego modules in the account.
func (r *CWSRegoModuleService) ListAutoPaging(ctx context.Context, params CWSRegoModuleListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[RegoModule] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes the module if no policy or library references its rules.
func (r *CWSRegoModuleService) Delete(ctx context.Context, regoModuleID string, body CWSRegoModuleDeleteParams, opts ...option.RequestOption) (res *CWSRegoModuleDeleteResponse, err error) {
	var env CWSRegoModuleDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if regoModuleID == "" {
		err = errors.New("missing required rego_module_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/rego-modules/%s", body.AccountID, regoModuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the library Rego module with the specified ID.
func (r *CWSRegoModuleService) Get(ctx context.Context, regoModuleID string, query CWSRegoModuleGetParams, opts ...option.RequestOption) (res *RegoModule, err error) {
	var env CWSRegoModuleGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if regoModuleID == "" {
		err = errors.New("missing required rego_module_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cws/rego-modules/%s", query.AccountID, regoModuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type RegoModule struct {
	// Server-assigned UUID
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Human-readable name; unique per account
	Name string `json:"name" api:"required"`
	// Dotted path derived from rego_source. Begins with lib. and is unique per
	// account.
	PackagePath string `json:"package_path" api:"required"`
	// Library Rego source with a package beginning with "lib." Maximum 16384 UTF-8
	// bytes.
	RegoSource string         `json:"rego_source" api:"required"`
	UpdatedAt  time.Time      `json:"updated_at" api:"required" format:"date-time"`
	JSON       regoModuleJSON `json:"-"`
}

// regoModuleJSON contains the JSON metadata for the struct [RegoModule]
type regoModuleJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	PackagePath apijson.Field
	RegoSource  apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegoModule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regoModuleJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleDeleteResponse struct {
	ID   string                          `json:"id" api:"required" format:"uuid"`
	JSON cwsRegoModuleDeleteResponseJSON `json:"-"`
}

// cwsRegoModuleDeleteResponseJSON contains the JSON metadata for the struct
// [CWSRegoModuleDeleteResponse]
type cwsRegoModuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleNewParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Name of the Rego module (unique per account)
	Name param.Field[string] `json:"name" api:"required"`
	// Library Rego source with a package beginning with "lib." Maximum 16384 UTF-8
	// bytes.
	RegoSource param.Field[string] `json:"rego_source" api:"required"`
}

func (r CWSRegoModuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSRegoModuleNewResponseEnvelope struct {
	Errors   []CWSRegoModuleNewResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSRegoModuleNewResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   RegoModule                                 `json:"result" api:"required"`
	Success  bool                                       `json:"success" api:"required"`
	JSON     cwsRegoModuleNewResponseEnvelopeJSON       `json:"-"`
}

// cwsRegoModuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSRegoModuleNewResponseEnvelope]
type cwsRegoModuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code" api:"required"`
	Message string                                       `json:"message" api:"required"`
	Source  CWSRegoModuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsRegoModuleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsRegoModuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSRegoModuleNewResponseEnvelopeErrors]
type cwsRegoModuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleNewResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer" api:"required"`
	JSON    cwsRegoModuleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsRegoModuleNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSRegoModuleNewResponseEnvelopeErrorsSource]
type cwsRegoModuleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	Source  CWSRegoModuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsRegoModuleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsRegoModuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSRegoModuleNewResponseEnvelopeMessages]
type cwsRegoModuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleNewResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    cwsRegoModuleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsRegoModuleNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSRegoModuleNewResponseEnvelopeMessagesSource]
type cwsRegoModuleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleUpdateParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Name of the Rego module (unique per account)
	Name param.Field[string] `json:"name" api:"required"`
	// Library Rego source with a package beginning with "lib." Maximum 16384 UTF-8
	// bytes.
	RegoSource param.Field[string] `json:"rego_source" api:"required"`
}

func (r CWSRegoModuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CWSRegoModuleUpdateResponseEnvelope struct {
	Errors   []CWSRegoModuleUpdateResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSRegoModuleUpdateResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   RegoModule                                    `json:"result" api:"required"`
	Success  bool                                          `json:"success" api:"required"`
	JSON     cwsRegoModuleUpdateResponseEnvelopeJSON       `json:"-"`
}

// cwsRegoModuleUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSRegoModuleUpdateResponseEnvelope]
type cwsRegoModuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	Source  CWSRegoModuleUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsRegoModuleUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsRegoModuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSRegoModuleUpdateResponseEnvelopeErrors]
type cwsRegoModuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer" api:"required"`
	JSON    cwsRegoModuleUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsRegoModuleUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSRegoModuleUpdateResponseEnvelopeErrorsSource]
type cwsRegoModuleUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	Source  CWSRegoModuleUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsRegoModuleUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsRegoModuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSRegoModuleUpdateResponseEnvelopeMessages]
type cwsRegoModuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    cwsRegoModuleUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsRegoModuleUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSRegoModuleUpdateResponseEnvelopeMessagesSource]
type cwsRegoModuleUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleListParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Sort direction (default: asc)
	Direction param.Field[CWSRegoModuleListParamsDirection] `query:"direction"`
	// Filter by Rego module name (case-insensitive substring match)
	Name param.Field[string] `query:"name"`
	// Field to order results by (default: created_at)
	Order param.Field[CWSRegoModuleListParamsOrder] `query:"order"`
	// Filter by package path (case-sensitive literal substring match)
	PackagePath param.Field[string] `query:"package_path"`
	// Number of items per page (1-100, default 20)
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [CWSRegoModuleListParams]'s query parameters as
// `url.Values`.
func (r CWSRegoModuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction (default: asc)
type CWSRegoModuleListParamsDirection string

const (
	CWSRegoModuleListParamsDirectionAsc  CWSRegoModuleListParamsDirection = "asc"
	CWSRegoModuleListParamsDirectionDesc CWSRegoModuleListParamsDirection = "desc"
)

func (r CWSRegoModuleListParamsDirection) IsKnown() bool {
	switch r {
	case CWSRegoModuleListParamsDirectionAsc, CWSRegoModuleListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by (default: created_at)
type CWSRegoModuleListParamsOrder string

const (
	CWSRegoModuleListParamsOrderName      CWSRegoModuleListParamsOrder = "name"
	CWSRegoModuleListParamsOrderCreatedAt CWSRegoModuleListParamsOrder = "created_at"
	CWSRegoModuleListParamsOrderUpdatedAt CWSRegoModuleListParamsOrder = "updated_at"
)

func (r CWSRegoModuleListParamsOrder) IsKnown() bool {
	switch r {
	case CWSRegoModuleListParamsOrderName, CWSRegoModuleListParamsOrderCreatedAt, CWSRegoModuleListParamsOrderUpdatedAt:
		return true
	}
	return false
}

type CWSRegoModuleDeleteParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSRegoModuleDeleteResponseEnvelope struct {
	Errors   []CWSRegoModuleDeleteResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSRegoModuleDeleteResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   CWSRegoModuleDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                          `json:"success" api:"required"`
	JSON     cwsRegoModuleDeleteResponseEnvelopeJSON       `json:"-"`
}

// cwsRegoModuleDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CWSRegoModuleDeleteResponseEnvelope]
type cwsRegoModuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code" api:"required"`
	Message string                                          `json:"message" api:"required"`
	Source  CWSRegoModuleDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsRegoModuleDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsRegoModuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSRegoModuleDeleteResponseEnvelopeErrors]
type cwsRegoModuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer" api:"required"`
	JSON    cwsRegoModuleDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsRegoModuleDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CWSRegoModuleDeleteResponseEnvelopeErrorsSource]
type cwsRegoModuleDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code" api:"required"`
	Message string                                            `json:"message" api:"required"`
	Source  CWSRegoModuleDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsRegoModuleDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsRegoModuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CWSRegoModuleDeleteResponseEnvelopeMessages]
type cwsRegoModuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer" api:"required"`
	JSON    cwsRegoModuleDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsRegoModuleDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSRegoModuleDeleteResponseEnvelopeMessagesSource]
type cwsRegoModuleDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleGetParams struct {
	// Cloudflare account ID (32-character hexadecimal string)
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type CWSRegoModuleGetResponseEnvelope struct {
	Errors   []CWSRegoModuleGetResponseEnvelopeErrors   `json:"errors" api:"required,nullable"`
	Messages []CWSRegoModuleGetResponseEnvelopeMessages `json:"messages" api:"required,nullable"`
	Result   RegoModule                                 `json:"result" api:"required"`
	Success  bool                                       `json:"success" api:"required"`
	JSON     cwsRegoModuleGetResponseEnvelopeJSON       `json:"-"`
}

// cwsRegoModuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CWSRegoModuleGetResponseEnvelope]
type cwsRegoModuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code" api:"required"`
	Message string                                       `json:"message" api:"required"`
	Source  CWSRegoModuleGetResponseEnvelopeErrorsSource `json:"source"`
	JSON    cwsRegoModuleGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// cwsRegoModuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CWSRegoModuleGetResponseEnvelopeErrors]
type cwsRegoModuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer" api:"required"`
	JSON    cwsRegoModuleGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// cwsRegoModuleGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [CWSRegoModuleGetResponseEnvelopeErrorsSource]
type cwsRegoModuleGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	Source  CWSRegoModuleGetResponseEnvelopeMessagesSource `json:"source"`
	JSON    cwsRegoModuleGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// cwsRegoModuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CWSRegoModuleGetResponseEnvelopeMessages]
type cwsRegoModuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CWSRegoModuleGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    cwsRegoModuleGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// cwsRegoModuleGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CWSRegoModuleGetResponseEnvelopeMessagesSource]
type cwsRegoModuleGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CWSRegoModuleGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cwsRegoModuleGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
